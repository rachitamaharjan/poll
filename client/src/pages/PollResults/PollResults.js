import React, { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { usePollContext } from '../../context/PollContext';

const PollResults = () => {
  const { id } = useParams();
  const { poll, fetchPollById, votePoll } = usePollContext();

  useEffect(() => {
    fetchPollById(id);
  }, [id]);

  const totalVotes = useMemo(() => {
    return poll?.options.reduce((acc, option) => acc + option.voteCount, 0) || 0;
  }, [poll]);

  const getVotePercentage = (voteCount) => {
    return totalVotes > 0 ? ((voteCount / totalVotes) * 100).toFixed(2) : 0;
  };

  if (!poll) return <p>Loading poll results...</p>;

  return (
    <>
    {poll ? (
      <div className="poll-results-container">
        <h2>{poll.question}</h2>
        <ul>
          {poll.options.map((option, index) => (
            <li key={index}>
              <div>
                <span>{option.text}</span>
                <span>{option.voteCount} votes ({getVotePercentage(option.voteCount)}%)</span>
              </div>
              <div style={{ background: "#ddd", height: "24px", width: "100%" }}>
                <div style={{ width: `${getVotePercentage(option.voteCount)}%`, background: "#4caf50", height: "100%" }}></div>
              </div>
            </li>
          ))}
        </ul>
      </div>
    ) : (
      <p>Loading Poll...</p>
    )}
  </>
  );
};

export default PollResults;