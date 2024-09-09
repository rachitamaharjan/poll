import { useState, useCallback, useEffect } from 'react';
import usePollStream from './usePollStream';

const usePollData = () => {
  const [polls, setPolls] = useState([]);
  const [poll, setPoll] = useState(null);
  const [loading, setLoading] = useState(false);  // For loading states
  const [error, setError] = useState(null);  // To handle errors

   // Create a new poll
   const createPoll = useCallback(async (pollData) => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch('http://localhost:8080/v1/api/polls/', {
        method: 'POST',
        body: JSON.stringify(pollData),
        headers: {
          'Content-Type': 'application/json',
        },
        mode: 'cors',
      });

      if (!response.ok) {
        const errorData = await response.text();
        console.error('Fetch error:', response.status, errorData);
        throw new Error('Failed to create poll');
      }
      const newPoll = await response.json();
      setPolls((prevPolls) => [...prevPolls, newPoll]);  // Add new poll to the list
    } catch (err) {
      console.error('Error creating poll:', err);
      setError('Failed to create poll');
    } finally {
      setLoading(false);
    }
  }, []);

  const fetchPolls = async () => {
    setLoading(true);
    setError(null);  // Reset error before fetching
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

  // Listen for live poll updates via SSE
  const livePollUpdate = usePollStream(poll?.id);

  // Update the current poll when new data comes in from SSE
  useEffect(() => {
    if (livePollUpdate) {
      debugger
      // Update the poll when new data comes in
      setPoll((prevPoll) => {
        if (prevPoll?.id === livePollUpdate.id) {
          return { ...prevPoll, ...livePollUpdate }; // Merge live updates into the poll
        }
        return prevPoll;
      });
    }
  }, [livePollUpdate]);

  return {
    createPoll,
    polls,
    poll,
    fetchPolls,
    fetchPollById,
    votePoll,
  };
};

export default usePollData;