import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getAnimals, deleteAnimal } from "../api/animalApi";
import { MainButton, useThemeParams } from "@vkruglikov/react-telegram-web-app";
import {
  Pagination,
  Card,
  Button,
  Text,
  Avatar,
  Title,
} from "@telegram-apps/telegram-ui";
import { Animal } from "../models/animal";

const AnimalList: React.FC = () => {
  const [, theme] = useThemeParams();
  const [animals, setAnimals] = useState<Animal[]>([]);
  const [pagination, setPagination] = useState({ page: 1, count: 10 });
  const navigate = useNavigate();

  const fetchAnimals = async () => {
    try {
      const data = await getAnimals();
      setAnimals(data);
    } catch {
      alert("Не удалось загрузить список животных.");
    }
  };

  const handleDelete = async (id?: string) => {
    if (window.confirm("Вы уверены, что хотите удалить это объявление?")) {
      try {
        await deleteAnimal(id);
        fetchAnimals();
      } catch {
        alert("Ошибка удаления объявления.");
      }
    }
  };
  useEffect(() => {
    fetchAnimals();
  }, []);

  const handleMainButtonClick = () => navigate("/add");

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

  // Расчет данных для текущей страницы
  const paginatedAnimals = animals.slice(
    (pagination.page - 1) * pagination.count,
    pagination.page * pagination.count
  );

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
      {paginatedAnimals.map((animal) => (
        <>
          <Card
            key={animal.id}
            style={{
              marginBottom: "12px",
              padding: "12px",
              background: theme.secondary_bg_color,
            }}
          >
            <div style={{ display: "flex", alignItems: "center" }}>
              <Avatar
                size={96}
                src={animal.photoURLs ? animal.photoURLs[0] : ""}
                alt={animal.name}
                style={{ marginRight: "16px" }}
              />
              <div>
                <Title>{animal.name}</Title>
                <Text>
                  Дата создания:{" "}
                  {new Date(animal.createdAt).toLocaleString("ru", {
                    year: "numeric",
                    month: "long",
                    day: "numeric",
                    hour: "2-digit",
                    minute: "2-digit",
                    second: "2-digit",
                  })}
                </Text>
              </div>
            </div>
            <div
              style={{
                marginTop: "12px",
                display: "flex",
                gap: "8px",
              }}
            >
              <Button onClick={() => navigate(`/edit/${animal.id}`)}>
                Редактировать
              </Button>
              <Button mode="bezeled" onClick={() => handleDelete(animal.id)}>
                Удалить
              </Button>
              <Button onClick={() => navigate(`/view/${animal.id}`)}>
                Подробнее
              </Button>
            </div>
          </Card>
        </>
      ))}
      <Pagination
        page={pagination.page}
        onChange={(_, page) => setPagination({ ...pagination, page })}
      />
    </div>
  );
};

export default AnimalList;
