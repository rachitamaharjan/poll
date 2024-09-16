import { useState, useCallback, useEffect } from 'react';
import axios from 'axios';
import usePollStream from './usePollStream';

const usePollData = () => {
  const [polls, setPolls] = useState([]);
  const [poll, setPoll] = useState(null);
  const [loading, setLoading] = useState(false);  // For loading states
  const [error, setError] = useState(null);  // To handle errors
  const [shareUrl, setShareUrl] = useState(''); // To store the share URL of the created poll

   // Create a new poll
   const createPoll = useCallback(async (pollData) => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.post('http://localhost:8080/v1/api/polls/', pollData, {
        headers: {
          'Content-Type': 'application/json',
        },
      });

      const newPoll = response.data;
      setPolls((prevPolls) => [...prevPolls, newPoll]);
      setShareUrl(newPoll.url);
    } catch (err) {
      console.error('Error creating poll:', err);
      setError('Failed to create poll');
    } finally {
      setLoading(false);
    }
  }, []);

  // Fetch all polls
  const fetchPolls = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get('/v1/api/polls');
      setPolls(response.data);
    } catch (err) {
      console.error('Error fetching polls:', err);
      setError('Failed to fetch polls');
    } finally {
      setLoading(false);
    }
  };

  // Fetch a poll by ID
  const fetchPollById = async (id) => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.get(`http://localhost:8080/v1/api/polls/${id}`, {
        headers: {
          'Content-Type': 'application/json',
        },
      });

      setPoll(response.data);
    } catch (err) {
      console.error('Error fetching poll:', err);
      setError('Failed to fetch poll');
    } finally {
      setLoading(false);
    }
  };

  // Vote on a poll
  const votePoll = async (id, optionIndex) => {
    try {
      await axios.post(`http://localhost:8080/v1/api/polls/${id}/vote`, { optionIndex }, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      fetchPollById(id); // Refresh poll data
    } catch (err) {
      console.error('Error voting on poll:', err);
      setError('Failed to vote on poll');
    }
  };

  // Listen for live poll updates via SSE
  const livePollUpdate = usePollStream(poll?.id);

  // Update the current poll when new data comes in from SSE
  useEffect(() => {
    if (livePollUpdate) {
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
    error,
    loading,
    shareUrl
  };
};

export default usePollData;