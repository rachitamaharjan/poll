import React from 'react';
import PropTypes from 'prop-types';
import { useTranslation } from 'react-i18next';

const ErrorBoundary = ({ children }) => {
  const { t } = useTranslation();
  const [hasError, setHasError] = React.useState(false);

  React.useEffect(() => {
    const handleError = (error) => {
      setHasError(true);
      console.error("ErrorBoundary caught an error", error);
    };

    window.addEventListener('error', handleError);
    return () => window.removeEventListener('error', handleError);
  }, []);

  if (hasError) {
    return (
      <div className="error-boundary">
        <h1>{t('errorBoundary.title')}</h1>
        <p>{t('errorBoundary.message')}</p>
        <button onClick={() => window.location.reload()}>{t('errorBoundary.refresh')}</button>
      </div>
    );
  }

  return children;
};

ErrorBoundary.propTypes = {
  children: PropTypes.node.isRequired,
};

export default ErrorBoundary;