import { Request, Response } from "express";
import { sendMail } from "../services/mail.service";
import { codeStore } from "../utils/codeStore";

export const sendVerificationEmail = async (req: Request, res: Response) => {
  const { email } = req.body;
  if (!email) return res.status(400).json({ message: "缺少邮箱地址" });

  const code = Math.floor(100000 + Math.random() * 900000).toString();
  codeStore.set(email, code);

  try {
    await sendMail(email, "邮箱验证", `您的验证码是：${code}`);
    res.status(200).json({ message: "验证码已发送，请检查邮箱" });
  } catch (err) {
    console.error(err);
    res.status(500).json({ message: "邮件发送失败" });
  }
};

export const verifyCode = (req: Request, res: Response) => {
  const { email, code } = req.body;
  if (!email || !code) return res.status(400).json({ message: "缺少参数" });

  const stored = codeStore.get(email);
  if (stored === code) {
    codeStore.delete(email);
    res.status(200).json({ message: "验证成功" });
  } else {
    res.status(400).json({ message: "验证码错误或已过期" });
  }
};
