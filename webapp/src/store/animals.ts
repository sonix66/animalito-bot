import {create} from "zustand";
import { Animal } from "../models/animal";

interface AnimalState {
  animals: Animal[];
  setAnimals: (animals: Animal[]) => void;
}

export const useAnimalStore = create<AnimalState>((set) => ({
  animals: [],
  setAnimals: (animals) => set({ animals }),
}));