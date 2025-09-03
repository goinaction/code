// cloud.js

export function saveToCloud(message, reply) {
  const data = {
    message: message,
    reply: reply,
    timestamp: new Date().toISOString()
  };

  // Simulated cloud storage
  console.log("Saving to cloud:", data);

  // Future: Connect to Google Drive or OneDrive API here
}
