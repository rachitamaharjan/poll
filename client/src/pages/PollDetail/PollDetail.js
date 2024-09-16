import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { usePollContext } from '../../context/PollContext';

import './PollDetail.css';

const PollDetail = () => {
  const { id } = useParams();
  const { poll, fetchPollById, votePoll, status, error } = usePollContext();
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
    votePoll(id, optionIndex)
  };

  const handleViewResults = () => {
    navigate(`/polls/${id}/results`);
  };

  const totalVotes = poll?.options?.reduce((sum, option) => sum + option.voteCount, 0) || 0;

  return (
    <div className="poll-detail-container">
      {loading ? (
        <p className="loading-message">Loading Poll...</p>
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
                      <span className="poll-option-votes"> - {option.voteCount} votes</span>
                    </li>
                  ))}
                </ul>
              </div>
              {totalVotes > 0 ? (
                <p className="total-votes">Total Votes: {totalVotes}</p>
              ) : (
                <p className="total-votes">Be the first to vote!</p>
              )}
              {status && <p className="vote-status-message">{status}</p>}
              {error && <p className="vote-error-message">{error}</p>}
              <button 
                className="view-results-button" 
                onClick={handleViewResults}
              >
                View Results
              </button>
            </>
          ) : (
            <p className="loading-message">Poll not found.</p>
          )}
        </div>
      )}
    </div>
  );
};

export default PollDetail;