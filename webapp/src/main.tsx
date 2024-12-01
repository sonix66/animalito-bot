import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import { WebAppProvider } from "@vkruglikov/react-telegram-web-app";
import { AppRoot } from "@telegram-apps/telegram-ui";
import "@telegram-apps/telegram-ui/dist/styles.css";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <AppRoot>
      <WebAppProvider
        options={{
          smoothButtonsTransition: true,
        }}
      >
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </WebAppProvider>
    </AppRoot>
  </React.StrictMode>
);
