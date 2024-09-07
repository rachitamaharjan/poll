import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { usePollContext } from '../../context/PollContext';


const PollList = () => {
  const { polls, fetchPolls } = usePollContext();

  useEffect(() => {
    fetchPolls();
  }, []);

  return (
    <div>
      <h1>Polls</h1>
      <ul>
        {polls?.map(poll => (
          <li key={poll.id}>
            <Link to={`/polls/${poll.id}`}>{poll.question}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PollList;