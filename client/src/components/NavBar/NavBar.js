import React from 'react';
import { Link } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import './NavBar.css';

const NavBar = ({ currentLanguage, onLanguageChange }) => {
  const { t } = useTranslation();

  return (
    <nav className="navbar">
      <ul className="nav-list">
        <li className="nav-item">
          <Link to="/" className="nav-link">{t('navbar.home')}</Link>
        </li>
        <li className="nav-item">
          <Link to="/polls/create" className="nav-link">{t('navbar.create_poll')}</Link>
        </li>
        <li className="nav-item selector">
          <button
            onClick={() => onLanguageChange('en')}
            className={`language-btn ${currentLanguage === 'en' ? 'active' : ''}`}
          >
            EN
          </button>
          <button
            onClick={() => onLanguageChange('np')}
            className={`language-btn ${currentLanguage === 'np' ? 'active' : ''}`}
          >
            NP
          </button>
        </li>
      </ul>
    </nav>
  );
};

export default NavBar;