import { Avatar, Button, Card, Title, Text } from "@telegram-apps/telegram-ui";
import React from "react";
import { Animal } from "../models/animal";
import { ThemeParams } from "@vkruglikov/react-telegram-web-app";
import { useNavigate } from "react-router-dom";

export type AnimalCardProps = {
  animal: Animal;
  theme: ThemeParams;
  onAnimalDelete: (id: string) => void;
};

export const AnimalCard: React.FC<AnimalCardProps> = ({ animal, theme, onAnimalDelete }) => {
  const navigate = useNavigate();
  return (
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
            {new Date(animal.createdAt!).toLocaleString("ru", {
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
        <Button mode="bezeled" onClick={() => onAnimalDelete(animal.id!)}>
          Удалить
        </Button>
        <Button onClick={() => navigate(`/view/${animal.id}`)}>
          Подробнее
        </Button>
      </div>
    </Card>
  );
};
