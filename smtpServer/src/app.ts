import express from "express";
import dotenv from "dotenv";
import emailRouter from "./routes/email.routes";
import cors from "cors";
dotenv.config();

const app = express();
app.use(cors({
  origin: "http://localhost:3000", // 前端地址
  methods: ["GET", "POST"],
  credentials: true, // 如果前端携带 cookies，需要设置 true
}));
app.use(express.json());
app.use("/email", emailRouter);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`✅ Server running`);
});
