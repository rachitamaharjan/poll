import React from 'react';
import './SuccessModal.css';

const SuccessModal = ({ isOpen, onClose, shareUrl }) => {
  if (!isOpen) return null;

  const handleOpenClick = () => {
    window.open(shareUrl, '_blank');
  };

  return (
    <div className="success-modal-overlay">
      <div className="success-modal-content">
        <h3>Poll Created Successfully!</h3>
        <p>Your poll has been created. Share it using the link below:</p>
        <div className="share-link-container">
          <input
            type="text"
            readOnly
            value={shareUrl}
            className="share-link"
          />
          <button onClick={handleOpenClick} className="open-link-btn">
            Open
          </button>
        </div>
        <button onClick={onClose} className="close-modal-btn">
          Close
        </button>
      </div>
    </div>
  );
};

export default SuccessModal;