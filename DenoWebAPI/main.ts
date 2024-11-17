/// <reference lib="deno.ns" />

import { sampleTodos } from './todo.ts';


// Create the handler function
const handler = (req: Request): Response => {
  // Only handle GET requests to `/todos`
  if (req.method === "GET" && new URL(req.url).pathname === "/todos") {
    return new Response(JSON.stringify(sampleTodos), {
      status: 200,
      headers: { "Content-Type": "application/json" },
    });
  }

  // Return 404 for other requests
  return new Response("Not Found", { status: 404 });
};

// Start the server
console.log("Server running on http://localhost:9090");
Deno.serve({ port: 9090 }, handler);
