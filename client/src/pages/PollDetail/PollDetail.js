import React, { useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { usePollContext } from '../../context/PollContext';

const PollDetail = () => {
  const { id } = useParams();
  const { poll, fetchPollById, votePoll } = usePollContext();

  useEffect(() => {
    fetchPollById(id);
  }, [id]);

  if (!poll) return <p>Loading poll...</p>;

  const handleVote = (optionIndex) => {
    votePoll(id, optionIndex);
  };

  return (
    <div>
      {poll ? (
        <div>
          <h1>{poll.question}</h1>
          <ul>
            {poll.options.map((option, index) => (
              <li key={index}>
                <button onClick={() => handleVote(index)}>{option.text}</button> - Votes: {option.voteCount}
              </li>
            ))}
          </ul>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default PollDetail;