import './App.css';
import { Routes, Route, Link } from 'react-router-dom';
import { PollProvider } from './context/PollContext';
import CreatePoll from './pages/CreatePoll/CreatePoll';
import Home from './pages/Home/Home';
import PollDetail from './pages/PollDetail/PollDetail';
import PollResults from './pages/PollResults/PollResults';


function App() {
  return (
    <PollProvider>
      <div>
        <nav>
            <ul>
                <li><Link to="/">Home</Link></li>
                <li><Link to="/polls/create/">Create Poll</Link></li>
            </ul>
        </nav>
        <Routes>
          <Route path="/" exact element={<Home />} />
          <Route path="/polls/create" element={<CreatePoll />} />
          <Route path="/polls/:id" element={<PollDetail />} />
          <Route path="/polls/:id/results" element={<PollResults />} />
        </Routes>
      </div>
  </PollProvider>
  );
}

export default App;
