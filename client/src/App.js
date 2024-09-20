import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import { PollProvider } from './context/PollContext';
import CreatePoll from './pages/CreatePoll/CreatePoll';
import Home from './pages/Home/Home';
import PollDetail from './pages/PollDetail/PollDetail';
import PollResults from './pages/PollResults/PollResults';
import NavBar from './components/NavBar/NavBar';
import { useTranslation } from 'react-i18next';

function App() {
  const { i18n } = useTranslation();
  const currentLanguage = i18n.language;

  const handleLanguageChange = (lng) => {
    i18n.changeLanguage(lng);
  };

  return (
    <PollProvider>
      <div className="app-container">
        <NavBar currentLanguage={currentLanguage} onLanguageChange={handleLanguageChange} />
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