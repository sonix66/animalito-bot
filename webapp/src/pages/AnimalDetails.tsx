import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  Button,
  Headline,
  Subheadline,
  Caption,
} from "@telegram-apps/telegram-ui";
import { getMockAnimalById, deleteMockAnimal } from "../api/animalApi";
import { BackButton, useThemeParams } from "@vkruglikov/react-telegram-web-app";
import { PhotoList } from "../components/PhotoList";
import { Animal } from "../models/animal";

const AnimalDetails: React.FC = () => {
  const [, theme] = useThemeParams();
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [animal, setAnimal] = useState<Animal | null>(null);

  useEffect(() => {
    if (id) {
      getMockAnimalById(id)
        .then(setAnimal)
        .catch(() => alert("Не удалось загрузить объявление."));
    }
  }, [id]);

  const handleDelete = async () => {
    try {
      await deleteMockAnimal(id!);
      alert("Объявление удалено.");
      navigate("/");
    } catch {
      alert("Ошибка удаления объявления.");
    }
  };

  const handleBackButtonClick = () => navigate("/");

  if (!animal) return null;

  return (
    <div
      style={{
        padding: "12px",
        gap: "8px",
        display: "flex",
        flexDirection: "column",
      }}
    >
      <BackButton onClick={handleBackButtonClick} />
      <PhotoList photos={animal.photoURLs} bgColor={theme.secondary_bg_color} />
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "12px",
          padding: "16px",
          background: theme.secondary_bg_color,
          borderRadius: "16px",
        }}
      >
        <Headline>{animal.name}</Headline>
        <Subheadline>{animal.type}</Subheadline>
        <Caption>{animal.description}</Caption>
        {animal.createdAt && (
          <Caption>
            Дата создания: {new Date(animal.createdAt).toLocaleDateString()}
          </Caption>
        )}
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            padding: 0,
          }}
        >
          <Button onClick={() => navigate(`/edit/${animal.id}`)}>
            Редактировать
          </Button>
          <Button mode="bezeled" onClick={handleDelete}>
            Удалить
          </Button>
        </div>
      </div>
    </div>
  );
};

export default AnimalDetails;