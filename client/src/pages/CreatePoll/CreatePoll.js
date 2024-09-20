import React, { useState, useCallback, Suspense } from 'react';
import { useNavigate } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { usePollContext } from '../../context/PollContext';
import ErrorBoundary from '../../components/ErrorBoundary/ErrorBoundary';
import SuccessModal from '../../components/SuccessModal/SuccessModal';

import './CreatePoll.css';

const CreatePoll = () => {
  const { createPoll, shareUrl } = usePollContext();
  const { t } = useTranslation();
  const [question, setQuestion] = useState('');
  const [allowMultipleVotes, setAllowMultipleVotes] = useState(false);
  const [options, setOptions] = useState([{ text: '' }]);
  const [error, setError] = useState(null);
  const [successModalOpen, setSuccessModalOpen] = useState(false);
  const navigate = useNavigate();

  // Performance optimization with useCallback
  const handleOptionChange = useCallback((index, event) => {
    const newOptions = [...options];
    newOptions[index].text = event.target.value;
    setOptions(newOptions);
  }, [options]);

  const addOption = useCallback(() => {
    setOptions([...options, { text: '' }]);
  }, [options]);

  const removeOption = useCallback((index) => {
    setOptions(options.filter((_, i) => i !== index));
  }, [options]);

  // Handling the form submission
  const handleSubmit = async (event) => {
    event.preventDefault();
    
    if (!question || options.some(option => option.text.trim() === '')) {
      setError(t('create_poll.error.empty_fields'));
      return;
    }

    try {
      await createPoll({ question, options, allowMultipleVotes }); // Call Context API method to create poll
      setSuccessModalOpen(true);
    } catch (error) {
      setError(t('create_poll.error.creation_failed'));
    }
  };

  const closeModal = () => {
    setSuccessModalOpen(false);
    navigate('/polls/create'); // Redirect to polls page after closing modal
  };

  return (
    <div className="create-poll-container">
      <h2>{t('create_poll.title')}</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="question">{t('create_poll.question_label')}:</label>
          <input
            type="text"
            id="question"
            value={question}
            onChange={(e) => setQuestion(e.target.value)}
            placeholder={t('create_poll.question_placeholder')}
            required
          />
        </div>

        <div className="form-group">
          <label>{t('create_poll.options_label')}:</label>
          {options.map((option, index) => (
            <div key={index} className="poll-option">
              <input
                type="text"
                value={option.text}
                onChange={(e) => handleOptionChange(index, e)}
                placeholder={`${t('create_poll.option')} ${index + 1}`}
                required
              />
              {options.length > 1 && (
                <button
                  type="button"
                  onClick={() => removeOption(index)}
                  className="remove-option-btn"
                >
                  {t('create_poll.remove_option')}
                </button>
              )}
            </div>
          ))}

          <button type="button" onClick={addOption} className="add-option-btn">
            {t('create_poll.add_option')}
          </button>
        </div>

        <div className="form-group">
          <input
            type="checkbox"
            id="allowMultipleVotes"
            checked={allowMultipleVotes}
            onChange={(e) => setAllowMultipleVotes(e.target.checked)}
          />
          <label htmlFor="allowMultipleVotes">{t('create_poll.allow_multiple_votes')}</label>
        </div>

        {error && <p className="error-message">{error}</p>}

        <button type="submit" className="submit-btn">
          {t('create_poll.submit_button')}
        </button>
      </form>

      <SuccessModal
        isOpen={successModalOpen}
        onClose={closeModal}
        shareUrl={shareUrl}
      />
    </div>
  );
};

// Wrapping CreatePoll with ErrorBoundary
const CreatePollWithErrorBoundary = () => (
  <ErrorBoundary>
    <Suspense fallback={<div>Loading...</div>}>
      <CreatePoll />
    </Suspense>
  </ErrorBoundary>
);

export default CreatePollWithErrorBoundary;