import { useEffect, useState } from 'react';

const usePollStream = (pollId) => {
  const [pollUpdates, setPollUpdates] = useState(null);

  useEffect(() => {
    if (!pollId) return; // Only initiate SSE if there's a pollId

    const eventSource = new EventSource(`http://localhost:8080/v1/api/polls/${pollId}/stream`);

    eventSource.onmessage = (event) => {
      debugger
      const pollData = JSON.parse(event.data);
      setPollUpdates(pollData);
    };

    eventSource.onerror = (err) => {
      console.error('SSE connection error:', err);
      eventSource.close();
    };

    return () => {
      eventSource.close();
    };
  }, [pollId]);

  return pollUpdates;
};

export default usePollStream;