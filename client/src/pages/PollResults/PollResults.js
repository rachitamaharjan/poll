import React, { useEffect, useMemo } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Pie } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { usePollContext } from '../../context/PollContext';
import './PollResults.css';

// Register required Chart.js components
ChartJS.register(ArcElement, Tooltip, Legend);

const PollResults = () => {
  const { id } = useParams();
  const { poll, fetchPollById, error, loading } = usePollContext();
  const navigate = useNavigate();

  useEffect(() => {
    fetchPollById(id);
  }, [id]);

  const totalVotes = useMemo(() => {
    return poll?.options.reduce((acc, option) => acc + option.voteCount, 0) || 0;
  }, [poll]);

  if (error) {
    return <p className="error-message">Error loading poll: {error}</p>;
  }

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

  if (loading) {
    return <p className="loading-message">Loading Poll...</p>;
  }

  return (
    <div className="poll-results-container">
      {poll ? (
        <>
          <h2 className="poll-question">{poll.question}</h2>
          <div className="poll-results-list">
            {poll.options?.map((option, index) => (
              <div key={index} className="poll-result-item">
                <div className="poll-option-text">
                  <span>{option.text}</span>
                  <span className="poll-option-stats">{option.voteCount} votes ({getVotePercentage(option.voteCount)}%)</span>
                </div>
                <div className="progress-bar-container">
                  <div 
                    className="progress-bar" 
                    style={{ width: `${getVotePercentage(option.voteCount)}%` }}
                  ></div>
                </div>
              </div>
            ))}
          </div>
          <p className="total-votes">Total Votes: {totalVotes}</p>
          <div className="pie-chart-container">
            <Pie data={pieData} />
          </div>
          <button className="back-to-poll-button" onClick={() => navigate(`/polls/${id}`)}>
            Back to Poll
          </button>
        </>
      ) : (
        <p className="loading-message">Poll not found.</p>
      )}
    </div>
  );
};

export default PollResults;