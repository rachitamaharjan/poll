import React from 'react';
import { useTranslation } from 'react-i18next';
import './SuccessModal.css';

const SuccessModal = ({ isOpen, onClose, shareUrl }) => {
  const { t } = useTranslation();

  if (!isOpen) return null;

  const handleOpenClick = () => {
    window.open(shareUrl, '_blank');
  };

  return (
    <div className="success-modal-overlay">
      <div className="success-modal-content">
        <h3>{t('success_modal.title')}</h3>
        <p>{t('success_modal.message')}</p>
        <div className="share-link-container">
          <input
            type="text"
            readOnly
            value={shareUrl}
            className="share-link"
          />
          <button onClick={handleOpenClick} className="open-link-btn">
            {t('success_modal.open_button')}
          </button>
        </div>
        <button onClick={onClose} className="close-modal-btn">
          {t('success_modal.close_button')}
        </button>
      </div>
    </div>
  );
};

export default SuccessModal;