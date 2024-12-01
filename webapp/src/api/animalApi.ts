import axios from "axios";
import { Animal } from "../models/animal";

const API_URL = "https://api.example.com/animals";

// Fetch all animals
export const getAnimals = async (): Promise<Animal[]> => {
  const { data } = await axios.get(API_URL);
  return data;
};

// Fetch a single animal by ID
export const getAnimalById = async (id: string): Promise<Animal> => {
  const { data } = await axios.get(`${API_URL}/${id}`);
  return data;
};

// Add a new animal
export const addAnimal = async (animal: Partial<Animal>): Promise<Animal> => {
  const { data } = await axios.post(API_URL, animal);
  return data;
};

// Update an existing animal
export const updateAnimal = async (id: string, animal: Partial<Animal>): Promise<Animal> => {
  const { data } = await axios.put(`${API_URL}/${id}`, animal);
  return data;
};

// Delete an animal
export const deleteAnimal = async (id: string): Promise<void> => {
  await axios.delete(`${API_URL}/${id}`);
};

// Fetch all animals
export const getMockAnimals = async (): Promise<Animal[]> => {
  return [
    {
      id: "1",
      photoURLs: ["https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"],
      type: "Dog",
      name: "Buddy",
      description: "A friendly dog.",
      createdAt: "2023-01-01",
    },
    {
      id: "2",
      photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
      type: "Cat",
      name: "Whiskers",
      description: "A playful cat.",
      createdAt: "2023-01-02",
    },
    {
      id: "1",
      photoURLs: ["https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"],
      type: "Dog",
      name: "Buddy",
      description: "A friendly dog.",
      createdAt: "2023-01-01",
    },
    {
      id: "2",
      photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
      type: "Cat",
      name: "Whiskers",
      description: "A playful cat.",
      createdAt: "2023-01-02",
    },
    {
      id: "1",
      photoURLs: ["https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"],
      type: "Dog",
      name: "Buddy",
      description: "A friendly dog.",
      createdAt: "2023-01-01",
    },
    {
      id: "2",
      photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
      type: "Cat",
      name: "Whiskers",
      description: "A playful cat.",
      createdAt: "2023-01-02",
    },
    {
      id: "1",
      photoURLs: ["https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"],
      type: "Dog",
      name: "Buddy",
      description: "A friendly dog.",
      createdAt: "2023-01-01",
    },
    {
      id: "2",
      photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
      type: "Cat",
      name: "Whiskers",
      description: "A playful cat.",
      createdAt: "2023-01-02",
    },
    {
      id: "1",
      photoURLs: ["https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"],
      type: "Dog",
      name: "Buddy",
      description: "A friendly dog.",
      createdAt: "2023-01-01",
    },
    {
      id: "2",
      photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
      type: "Cat",
      name: "Whiskers",
      description: "A playful cat.",
      createdAt: "2023-01-02",
    },
    {
      id: "1",
      photoURLs: ["https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"],
      type: "Dog",
      name: "Buddy",
      description: "A friendly dog.",
      createdAt: "2023-01-01",
    },
    {
      id: "2",
      photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
      type: "Cat",
      name: "Whiskers",
      description: "A playful cat.",
      createdAt: "2023-01-02",
    },
  ];
};

// Fetch a single animal by ID
export const getMockAnimalById = async (id: string): Promise<Animal> => {
  return {
    id,
    photoURLs: [
      "https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png",
      "https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png",
      "https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png",
      "https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png",
      "https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png",
      // "https://static.insales-cdn.com/files/1/2278/35317990/original/mceu_288641621718726728282-1718726728309.png"
    ],
    type: "Dog",
    name: "Buddy",
    description: "A friendly dog.",
    createdAt: "2023-01-01",
  };
};

// Add a new animal
export const addMockAnimal = async (animal: Partial<Animal>): Promise<Animal> => {
  return {
    id: "3",
    photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
    type: "Cat",
    name: "Whiskers",
    description: "A playful cat.",
    createdAt: "2023-01-02",
  }
};

// Update an existing animal
export const updateMockAnimal = async (id: string, animal: Partial<Animal>): Promise<Animal> => {
  return {
    id: "3",
    photoURLs: ["https://avatars.mds.yandex.net/get-altay/11400795/2a0000018dad68dcbd98f232ecca031975a3/L_height"],
    type: "Cat",
    name: "Whiskers",
    description: "A playful cat.",
    createdAt: "2023-01-02",
  }
};

// Delete an animal
export const deleteMockAnimal = async (id?: string): Promise<void> => {};
