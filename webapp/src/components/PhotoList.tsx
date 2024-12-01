import { List, Image } from "@telegram-apps/telegram-ui";
import React from "react";

export type PhotoListProps = {
  photos: string[];
  bgColor?: string;
};

export const PhotoList: React.FC<PhotoListProps> = ({ photos, bgColor }) => {
  if (!photos || photos.length < 1) {
    return null;
  }

  return (
    <List
      style={{
        display: "flex",
        justifyContent: "center",
        gap: "12px",
        width: "100%",
        flexWrap: "wrap",
        background: bgColor,
        borderRadius: "16px",
        padding: "16px",
      }}
    >
      {photos.map((photo, index) => (
        <Image key={index} src={photo} size={96} style={{ margin: 0 }} />
      ))}
    </List>
  );
};
