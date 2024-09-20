import React from 'react';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import './Home.css';

const HomePage = () => {
  const { t } = useTranslation();

  return (
    <div className="home-page-container">
      <h1 className="welcome-message">{t('home_page.welcome_message')}</h1>
      <Link to="/polls/create">
        <button className="create-poll-button">{t('home_page.create_poll')}</button>
      </Link>
    </div>
  );
};

export default HomePage;