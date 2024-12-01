import React from "react";
import { Routes, Route } from "react-router-dom";
import AnimalList from "./pages/AnimalList";
import AnimalForm from "./pages/AnimalForm";
import AnimalDetails from "./pages/AnimalDetails";
import { useThemeParams } from "@vkruglikov/react-telegram-web-app";

const App: React.FC = () => {
  const [colorScheme, themeParams] = useThemeParams();

  console.log(colorScheme === "dark");
  console.log({
    text_color: themeParams.text_color,
    button_color: themeParams.button_color,
    bg_color: themeParams.bg_color,
  });
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
