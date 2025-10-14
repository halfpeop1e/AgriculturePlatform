import { Router } from "express";
import { sendVerificationEmail, verifyCode } from "../controllers/email.controller";

const router = Router();

router.post("/sendCode", sendVerificationEmail);
router.post("/verify", verifyCode);

export default router;
