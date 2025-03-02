import React, { useState, useEffect } from 'react';
import './App.css';

const App = () => {
  const [numbers, setNumbers] = useState([]);
  const [currentNumber, setCurrentNumber] = useState('');
  const [target, setTarget] = useState('');
  const [calculationResult, setCalculationResult] = useState(null);

  useEffect(() => {
    setNumbers([250, 500, 1000, 2000, 5000]);
  }, []);

  const handleAddNumber = () => {
    if (currentNumber) {
      setNumbers([...numbers, Number(currentNumber)]);
      setCurrentNumber('');
    }
  };

  const handleDeleteNumber = (index) => {
    setNumbers(numbers.filter((_, i) => i !== index));
  };

  const handleSubmit = async () => {
    const requestData = {
      numbers: numbers,
      target: parseInt(target),
    };
    try {
      const response = await fetch('http://localhost:8080/calculate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestData),
      });
      const data = await response.json();
      setCalculationResult(data);
    } catch (error) {
      console.error('Error sending numbers to backend', error);
    }
  };

  return (
    <div className="container">
      <h2 className="title">Pack Sizes & Calculation</h2>

      <div className="input-container">
        <input
          type="number"
          className="input-field"
          value={currentNumber}
          onChange={(e) => setCurrentNumber(e.target.value)}
          placeholder="Add a pack size"
        />
        <button className="add-btn" onClick={handleAddNumber}>Add</button>
      </div>

      <div className="list-container">
        {numbers.map((num, index) => (
          <div className="number-item" key={index}>
            <span className="number">{num}</span>
            <span className="delete-btn" onClick={() => handleDeleteNumber(index)}>✖</span>
          </div>
        ))}
      </div>

      {/* Target Value Input */}
      <div className="input-container">
        <input
          type="number"
          className="input-field"
          value={target}
          onChange={(e) => setTarget(e.target.value)}
          placeholder="Enter the order size"
        />
      </div>

      {/* Submit Button */}
      <div className="submit-container">
        <button className="submit-btn" onClick={handleSubmit}>Calculate packs</button>
      </div>

      {/* Display the result */}
      {calculationResult && typeof calculationResult === 'object' && (
        <div className="result-container">
          <h3>Calculation Result:</h3>
          <ul className="result-list">
            {Object.entries(calculationResult).map(([number, count]) => (
              <li key={number} className="result-item">
                Pack Size: <strong>{number}</strong> → Used: <strong>{count} times</strong>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default App;
