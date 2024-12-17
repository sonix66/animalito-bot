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
import { BackButton, useThemeParams } from "@vkruglikov/react-telegram-web-app";
import { PhotoList } from "../components/PhotoList";
import { Animal } from "../models/animal";

const AnimalForm: React.FC = () => {
  const [photoPreview, setPhotoPreview] = useState<string | null>(null);
  const [, theme] = useThemeParams();
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    photo: null as File | null,
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
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];
      setFormData((prev) => ({
        ...prev,
        photo: file,
      }));
      setPhotoPreview(URL.createObjectURL(file)); // Создаем превью
    }
  };

  const handleSubmit = async () => {
    try {
      const dataToSend = new FormData();
      if (formData.photo) {
        dataToSend.append("photo", formData.photo); // Один файл
      }
      dataToSend.append("type", formData.type);
      dataToSend.append("name", formData.name);
      dataToSend.append("description", formData.description);

      if (id) {
        await updateAnimal(id, dataToSend as Partial<Animal>);
      } else {
        await addAnimal(dataToSend as Partial<Animal>);
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
          <>
            <FileInput
              name="photo"
              onChange={handleFileChange}
              accept="image/*"
              label={photoPreview ? "Изменить фото" : "Прикрепить фото"}
            />
            {photoPreview && (
              <div style={{ textAlign: "center", marginBottom: "12px" }}>
                <img
                  src={photoPreview}
                  alt="Preview"
                  style={{
                    maxWidth: "100%",
                    maxHeight: "200px",
                    borderRadius: "8px",
                  }}
                />
              </div>
            )}
          </>
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
