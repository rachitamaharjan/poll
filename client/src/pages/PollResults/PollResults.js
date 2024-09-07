import React, { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { Pie } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { usePollContext } from '../../context/PollContext';
import './PollResults.css';

// Register required Chart.js components
ChartJS.register(ArcElement, Tooltip, Legend);

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


  const pieData = {
    labels: poll?.options.map(option => option.text),
    datasets: [
      {
        label: 'Votes',
        data: poll?.options.map(option => option.voteCount),
        backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#4CAF50', '#FF9F40'],
        hoverBackgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#4CAF50', '#FF9F40'],
      },
    ],
  };

  return (
    <>
    {poll ? (
      <div className="poll-results-container">
        <h2>{poll.question}</h2>
        <ul>
          {poll.options?.map((option, index) => (
            <li key={index}>
              <div>
                <span>{option?.text}</span>
                <span>{option?.voteCount} votes ({getVotePercentage(option.voteCount)}%)</span>
              </div>
              <div style={{ background: "#ddd", height: "24px", width: "100%" }}>
                <div style={{ width: `${getVotePercentage(option.voteCount)}%`, background: "#4caf50", height: "100%" }}></div>
              </div>
            </li>
          ))}
        </ul>
        <p>Total Votes: {totalVotes}</p>
        {/* Pie Chart */}
        <div style={{ maxWidth: '400px', margin: 'auto' }}>
          <Pie data={pieData} />
        </div>
      </div>
    ) : (
      <p>Loading Poll...</p>
    )}
  </>
  );
};

export default PollResults;