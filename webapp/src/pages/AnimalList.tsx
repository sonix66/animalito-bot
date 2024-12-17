import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getAnimals, deleteAnimal, getAnimalsCount } from "../api/animalApi";
import { MainButton, useThemeParams } from "@vkruglikov/react-telegram-web-app";
import {
  Pagination,
  Text,
  Spinner
} from "@telegram-apps/telegram-ui";
import { useAnimalStore } from "../store/animals";
import { useQuery } from "@tanstack/react-query";
import { AnimalCard } from "../components/AnimalCard";

const AnimalList: React.FC = () => {
  const [, theme] = useThemeParams();
  const [pagination, setPagination] = useState({ page: 1, count: 10 });
  const [animalsCount, setAnimalsCount] = useState(0);
  const navigate = useNavigate();
  const { animals, setAnimals } = useAnimalStore();

  const { data, isLoading, error } = useQuery({
    queryKey: ["animals", pagination],
    queryFn: () => getAnimals(10, (pagination.page - 1) * 10),
  });
  const { data: dataCount, isLoading: isLoadingCount, error: errorCount } = useQuery({
    queryKey: ["animalsCount"],
    queryFn: () => getAnimalsCount(),
  })

  useEffect(() => {
    if (data) setAnimals(data);
    else setAnimals([]);
  }, [data, setAnimals]);

  useEffect(() => {
    if (dataCount) setAnimalsCount(dataCount.count);
    else setAnimalsCount(0) ;
  }, [dataCount, setAnimals]);

  const handleAnimalDelete = async (id?: string) => {
    if (window.confirm("Вы уверены, что хотите удалить это объявление?")) {
      try {
        await deleteAnimal(id!);
        setAnimals(animals.filter((animal) => animal.id !== id));
      } catch {
        alert("Ошибка удаления объявления.");
      }
    }
  };

  const handleMainButtonClick = () => navigate("/add");

  if (isLoading || isLoadingCount) return (
    <div style={{ display: "flex", justifyContent: "center", alignItems: "center", height: "100vh" }}>
      <Spinner size="l"/>
    </div>
  );
  if (error || errorCount) return <Text>Ошибка загрузки.</Text>;

  if (!animals || animals.length === 0) {
    return (
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          padding: "12px",
          gap: "4px",
        }}
      >
        <MainButton
          text="Добавить объявление"
          onClick={handleMainButtonClick}
        />
        <Text>Список животных пуст.</Text>
      </div>
    );
  }

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        padding: "12px",
        gap: "4px",
      }}
    >
      <MainButton text="Добавить объявление" onClick={handleMainButtonClick} />
      {data?.map((animal) => <AnimalCard key={animal.id} animal={animal} theme={theme} onAnimalDelete={handleAnimalDelete} />)}
      <Pagination
        page={pagination.page}
        onChange={(_, page) => setPagination({ ...pagination, page })}
        count={Math.ceil(animalsCount / pagination.count)}
      />
    </div>
  );
};

export default AnimalList;
