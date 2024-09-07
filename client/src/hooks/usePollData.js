import { useState } from 'react';

const usePollData = () => {
  const [polls, setPolls] = useState([]);
  const [poll, setPoll] = useState(null);

  const fetchPolls = async () => {
    const response = await fetch('/v1/api/polls');
    const data = await response.json();
    setPolls(data);
  };

  const fetchPollById = async (id) => {
    try {
      const response = await fetch(`http://localhost:8080/v1/api/polls/${id}`,
        {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      if (!response.ok) {
        throw new Error('Failed to fetch poll');
      }
      const data = await response.json();
      setPoll(data);
    } catch (error) {
      console.error('Error fetching poll:', error);
    }
  };

  const votePoll = async (id, optionIndex) => {
    await fetch(`http://localhost:8080/v1/api/polls/${id}/vote`, {
      method: 'POST',
      body: JSON.stringify({ optionIndex }),
      headers: {
        'Content-Type': 'application/json',
      },
    });
    fetchPollById(id); // Refresh poll data
  };

  return {
    polls,
    poll,
    fetchPolls,
    fetchPollById,
    votePoll,
  };
};

export default usePollData;