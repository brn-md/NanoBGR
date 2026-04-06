import React, { useState, useEffect } from 'react';
import { Upload, X, Check, Loader2, Download } from 'lucide-react';

function App() {
  const [file, setFile] = useState(null);
  const [status, setStatus] = useState(null);
  const [taskId, setTaskId] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleUpload = async () => {
    if (!file) return;
    setLoading(true);
    
    const formData = new FormData();
    formData.append('image', file);

    try {
      const response = await fetch('http://localhost:3000/upload', {
        method: 'POST',
        body: formData,
      });
      const data = await response.json();
      setTaskId(data.id);
      setStatus('processing');
    } catch (err) {
      alert('Error uploading image');
      setLoading(false);
    }
  };

  useEffect(() => {
    let interval;
    if (taskId && status === 'processing') {
      interval = setInterval(async () => {
        try {
          const res = await fetch(`http://localhost:3000/status/${taskId}`);
          const data = await res.json();
          if (data.status === 'done') {
            setStatus('done');
            setLoading(false);
            clearInterval(interval);
          }
        } catch (err) {
          console.error('Error checking status', err);
        }
      }, 2000);
    }
    return () => clearInterval(interval);
  }, [taskId, status]);

  return (
    <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', padding: '2rem' }}>
      <h1>NanoBGR 🌪️</h1>
      <p>AI Background Removal Project</p>

      {!taskId && (
        <div style={{ border: '2px dashed #334155', padding: '3rem', borderRadius: '1rem', textAlign: 'center' }}>
          <input type="file" onChange={(e) => setFile(e.target.files[0])} style={{ marginBottom: '1rem' }} />
          <button onClick={handleUpload} disabled={!file || loading}>
            {loading ? <Loader2 className="animate-spin" /> : <Upload />} Process High-Res Imagem
          </button>
        </div>
      )}

      {status === 'processing' && (
        <div style={{ textAlign: 'center' }}>
          <Loader2 style={{ animation: 'spin 2s linear infinite' }} size={48} />
          <p>Remove Background IA is Working...</p>
        </div>
      )}

      {status === 'done' && (
        <div style={{ textAlign: 'center' }}>
          <h3>✨ Background Removed!</h3>
          <p>Check the results on Minio Console or check logic for download.</p>
          <button onClick={() => { setTaskId(null); setStatus(null); setFile(null); }}>
             Upload Another One
          </button>
        </div>
      )}
    </div>
  );
}

export default App;
