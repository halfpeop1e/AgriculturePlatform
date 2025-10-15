import nodemailer from "nodemailer";
import dotenv from "dotenv";

dotenv.config();

const transporter = nodemailer.createTransport({
  host: process.env.SMTP_HOST,
  port: Number(process.env.SMTP_PORT),
  secure: true,
  auth: {
    user: process.env.SMTP_USER,
    pass: process.env.SMTP_PASS,
  },
});

export const sendMail = async (to: string, text: string) => {
  await transporter.sendMail({
    from: `"邮箱验证系统" <${process.env.SMTP_USER}>`,
    to,
    subject:"验证您的邮箱",
    text,
  });
};
