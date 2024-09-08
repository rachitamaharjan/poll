import './App.css';
import { Routes, Route } from 'react-router-dom';
import { PollProvider } from './context/PollContext';
import CreatePoll from './pages/CreatePoll/CreatePoll';
import PollList from './pages/PollList/PollList';
import PollDetail from './pages/PollDetail/PollDetail';
import PollResults from './pages/PollResults/PollResults';


function App() {
  return (
    <PollProvider>
      <div>
        <Routes>
          <Route path="/polls/create" element={<CreatePoll />} />
          <Route path="/polls" element={<PollList />} />
          <Route path="/polls/:id" element={<PollDetail />} />
          <Route path="/polls/:id/results" element={<PollResults />} />
        </Routes>
      </div>
  </PollProvider>
  );
}

export default App;
