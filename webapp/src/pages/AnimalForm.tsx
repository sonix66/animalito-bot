import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  Button,
  FileInput,
  Input,
  Textarea,
  Typography,
} from "@telegram-apps/telegram-ui";
import { addAnimal, updateAnimal, getAnimalById } from "../api/animalApi";
import {
  BackButton,
  MainButton,
  useThemeParams,
} from "@vkruglikov/react-telegram-web-app";
import { PhotoList } from "../components/PhotoList";
import { Animal } from "../models/animal";

const AnimalForm: React.FC = () => {
  const [, theme] = useThemeParams();
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    photos: [] as File[], // Здесь храним загруженные файлы
    type: "",
    name: "",
    description: "",
  });
  const [photos, setPhotos] = useState<string[]>([]);

  useEffect(() => {
    if (id) {
      getAnimalById(id).then((animal: Animal) => {
        setFormData((prev) => ({ ...prev, ...animal, photos: [] }));
        setPhotos(animal.photoURLs);
      });
    }
  }, [id]);

  const handleInputChange = (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = event.target;
    console.log({ name, value });
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files) {
      setFormData((prev) => ({
        ...prev,
        photos: Array.from(event.target.files), // Преобразуем FileList в массив
      }));
    }
  };

  const handleSubmit = async () => {
    try {
      const dataToSend = new FormData();
      formData.photos.forEach((photo, index) => {
        dataToSend.append(`photo_${index}`, photo);
      });
      dataToSend.append("type", formData.type);
      dataToSend.append("name", formData.name);
      dataToSend.append("description", formData.description);

      if (id) {
        await updateAnimal(id, dataToSend);
      } else {
        await addAnimal(dataToSend);
      }

      navigate("/");
    } catch {
      alert("Ошибка сохранения объявления");
    }
  };

  const handleBackButtonClick = () => navigate("/");

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
      <PhotoList photos={photos} bgColor={theme.secondary_bg_color} />
      <form
        onSubmit={handleSubmit}
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "12px",
          padding: "16px",
          background: theme.secondary_bg_color,
          borderRadius: "16px",
        }}
        encType="multipart/form-data"
      >
        {!id && (
          <FileInput
            name="photos"
            onChange={handleFileChange}
            accept="image/*"
            label={id ? "Обновить фото" : "Прикрепить фото"}
            multiple
          />
        )}
        <Typography>Заголовок объявления</Typography>
        <Input
          name="name"
          value={formData.name}
          onChange={handleInputChange}
          required
        />
        <Typography>Описание объявления</Typography>
        <Textarea
          name="description"
          value={formData.description}
          onChange={handleInputChange}
          required
        />
        <Button onClick={handleSubmit}>{id ? "Обновить" : "Сохранить"}</Button>
      </form>
    </div>
  );
};

export default AnimalForm;
