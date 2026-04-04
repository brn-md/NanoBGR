import cv2
import numpy as np
from rembg import remove, new_session
from PIL import Image

session = new_session('u2net')

def process_image(image_bytes: bytes) -> bytes:
    # 1. Decode original high-res image
    nparr = np.frombuffer(image_bytes, np.uint8)
    original_img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
    
    if original_img is None:
        raise ValueError("Failed to decode image")

    h_orig, w_orig = original_img.shape[:2]

    # 2. Downscale for proxy (e.g., max dimension 1024)
    max_dim = 1024
    scale = min(max_dim / w_orig, max_dim / h_orig)
    
    if scale < 1.0:
        new_w, new_h = int(w_orig * scale), int(h_orig * scale)
        proxy_img = cv2.resize(original_img, (new_w, new_h), interpolation=cv2.INTER_AREA)
    else:
        proxy_img = original_img.copy()

    # Convert proxy to PIL for rembg
    proxy_pil = Image.fromarray(cv2.cvtColor(proxy_img, cv2.COLOR_BGR2RGB))

    # 3. Generate mask using rembg on proxy
    result_pil = remove(proxy_pil, session=session, output_format="RGBA")
    
    # Extract alpha channel (the mask)
    result_np = np.array(result_pil)
    proxy_mask = result_np[:, :, 3]

    # 4. Upscale mask to original dimensions
    up_mask = cv2.resize(proxy_mask, (w_orig, h_orig), interpolation=cv2.INTER_LINEAR)

    # Optional: Apply some smoothing/blur to the upscaled mask
    up_mask = cv2.GaussianBlur(up_mask, (5, 5), 0)

    # 5. Apply the mask to the original image
    b, g, r = cv2.split(original_img)
    final_img = cv2.merge([b, g, r, up_mask])

    # Encode to PNG bytes
    success, buffer = cv2.imencode('.png', final_img)
    if not success:
        raise ValueError("Failed to encode final image")

    return buffer.tobytes()
