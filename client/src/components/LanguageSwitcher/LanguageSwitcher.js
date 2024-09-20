import React from 'react';
import { useTranslation } from 'react-i18next';

import './LanguageSwitcher.css'

const LanguageSwitcher = () => {
  const { i18n } = useTranslation();

  const switchLanguage = (lng) => {
    i18n.changeLanguage(lng);
  };

  return (
    <div className="language-switcher">
      <button onClick={() => switchLanguage('en')}>English</button>
      <button onClick={() => switchLanguage('np')}>Nepali</button>
    </div>
  );
};

export default LanguageSwitcher;