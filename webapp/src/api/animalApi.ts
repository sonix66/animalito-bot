import axios from "axios";
import { Animal } from "../models/animal";

const API_URL = "https://0m8gbg09-8080.euw.devtunnels.ms/api/animals";

// Fetch all animals
export const getAnimals = async (count: number, offset: number): Promise<Animal[]> => {
  const { data } = await axios.get(API_URL, {
    params: {count, offset}
  });
  return data;
};

// Fetch a single animal by ID
export const getAnimalById = async (id: string): Promise<Animal> => {
  const { data } = await axios.get(`${API_URL}/${id}`);
  return data;
};

// Fetch a single animal by ID
export const getAnimalsCount = async (): Promise<{ count: number }> => {
  const { data } = await axios.get(`${API_URL}/count`);
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
