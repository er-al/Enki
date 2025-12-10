import type React from 'react';
import './App.css';
import TypingTest from './components/TypingTest';

const App: React.FC = () => {
  return (
    <div className="content">
      <h1>Speed Typing Test</h1>
      <p>Test your typing speed! Type the words below as fast as you can.</p>
      <TypingTest />
    </div>
  );
};

export default App;
