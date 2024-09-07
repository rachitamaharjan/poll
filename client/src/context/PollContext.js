import React, { createContext, useState, useContext } from 'react';
import usePollData from '../hooks/usePollData';

const PollContext = createContext();

export const PollProvider = ({ children }) => {
  const pollData = usePollData();

  return (
    <PollContext.Provider value={pollData}>
      {children}
    </PollContext.Provider>
  );
};

export const usePollContext = () => {
  return useContext(PollContext);
};