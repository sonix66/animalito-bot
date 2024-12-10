import React from "react";
import { Routes, Route } from "react-router-dom";
import AnimalList from "./pages/AnimalList";
import AnimalForm from "./pages/AnimalForm";
import AnimalDetails from "./pages/AnimalDetails";

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<AnimalList />} />
      <Route path="/add" element={<AnimalForm />} />
      <Route path="/edit/:id" element={<AnimalForm />} />
      <Route path="/view/:id" element={<AnimalDetails />} />
    </Routes>
  );
};

export default App;
