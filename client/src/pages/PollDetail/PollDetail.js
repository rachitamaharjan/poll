import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { usePollContext } from '../../context/PollContext';
import { useTranslation } from 'react-i18next';
import './PollDetail.css';

const PollDetail = () => {
  const { id } = useParams();
  const { poll, fetchPollById, votePoll, status, error } = usePollContext();
  const { t } = useTranslation();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadPoll = async () => {
      try {
        await fetchPollById(id);
        setLoading(false);
      } catch (error) {
        setLoading(false);
      }
    };

    loadPoll();
  }, [id]);

  const handleVote = (optionIndex) => {
    votePoll(id, optionIndex);
  };

  const handleViewResults = () => {
    navigate(`/polls/${id}/results`);
  };

  const totalVotes = poll?.options?.reduce((sum, option) => sum + option.voteCount, 0) || 0;

  return (
    <div className="poll-detail-container">
      {loading ? (
        <p className="loading-message">{t('poll_detail.loading')}</p>
      ) : (
        <div className="poll-detail-content">
          {poll ? (
            <>
              <h2 className="poll-question">{poll.question}</h2>
              <div className="poll-options">
                <ul>
                  {poll.options?.map((option, index) => (
                    <li key={index} className="poll-option-item">
                      <button 
                        className="poll-option-button" 
                        onClick={() => handleVote(index)}
                      >
                        {option.text}
                      </button>
                      <span className="poll-option-votes"> - {option.voteCount} {t('poll_detail.votes')}</span>
                    </li>
                  ))}
                </ul>
              </div>
              {totalVotes > 0 ? (
                <p className="total-votes">{t('poll_detail.total_votes', { count: totalVotes })}</p>
              ) : (
                <p className="total-votes">{t('poll_detail.be_the_first')}</p>
              )}
              {status && <p className="vote-status-message">{status}</p>}
              {error && <p className="vote-error-message">{error}</p>}
              <button 
                className="view-results-button" 
                onClick={handleViewResults}
              >
                {t('poll_detail.view_results')}
              </button>
            </>
          ) : (
            <p className="loading-message">{t('poll_detail.not_found')}</p>
          )}
        </div>
      )}
    </div>
  );
};

export default PollDetail;