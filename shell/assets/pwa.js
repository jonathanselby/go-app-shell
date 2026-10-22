"use strict";
(() => {
  if (!("serviceWorker" in navigator)) return;
  const script = document.currentScript;
  const workerURL = script && script.dataset ? script.dataset.workerUrl : "";
  if (!workerURL) return;
  window.addEventListener("load", () => navigator.serviceWorker.register(workerURL).catch(() => {}), { once: true });
})();
