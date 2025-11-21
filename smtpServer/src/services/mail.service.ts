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

export const sendMail = async (to: string, text: string,code:string,actionName:string) => {
  await transporter.sendMail({
    from: `"邮箱验证系统" <${process.env.SMTP_USER}>`,
    to,
    subject:"验证您的邮箱",
    text,
    html: `
     <div style="
  background: linear-gradient(135deg, #d9e8ff 0%, #f5e8ff 100%);
  padding: 30px 0;
  font-family: Arial, Helvetica, sans-serif;
">
  <div style="
    max-width: 420px;
    margin: 0 auto;
    padding: 0;
    
    /* 毛玻璃半透明效果 */
    background: rgba(255, 255, 255, 0.55);
    backdrop-filter: blur(14px);
    -webkit-backdrop-filter: blur(14px);

    border-radius: 18px;
    border: 1px solid rgba(255, 255, 255, 0.35);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.12);
  ">
      
    <!-- 顶部品牌栏 -->
    <div style="
      padding: 20px; 
      text-align: center;
      display: flex;
      align-items: center;
      gap: 12px;
      justify-content: center;
      border-bottom: 1px solid rgba(255,255,255,0.4);
      background: rgba(255,255,255,0.4);
      backdrop-filter: blur(6px);
    ">
      <img src="http://47.94.106.89:8080/static/files/webIcon.jpg" alt="Logo" style="height:42px;width:42px; border-radius:8px;">
      <h2 style="margin:0; font-size:20px; font-weight:600;">农产品融销一体平台</h2>
    </div>

    <!-- 内容主体 -->
    <div style="padding: 32px 40px;">
      
      <h2 style="color:#2a2a2a; margin:0 0 16px; font-size:22px;">
        您好！
      </h2>

      <p style="color:#333; font-size:15px; line-height:1.7;">
        您正在进行 <strong>${actionName}</strong> 操作。  
        为保障账户安全，请使用下方验证码完成验证：
      </p>

      <!-- 验证码 -->
      <div style="
        margin: 24px 0; 
        text-align:center; 
        font-size:32px; 
        letter-spacing:6px; 
        font-weight:bold; 
        color:#1e63d8;
        padding: 16px 0;
        border-radius: 10px;

        background: rgba(240, 245, 255, 0.7);
        border: 1px dashed #1e63d8;
        backdrop-filter: blur(3px);
      ">
        ${code}
      </div>  

      <p style="color:#555; font-size:14px; line-height:1.7; margin-top:24px;">
        如果您未请求此操作，请忽略此邮件。  
        该验证码将在 <strong>10 分钟内</strong> 失效。
      </p>

      <!-- 分割线 -->
      <div style="height:1px; background:rgba(255,255,255,0.4); margin:32px 0;"></div>

      <!-- 页脚 -->
      <p style="color:#666; font-size:12px; text-align:center; line-height:1.6;">
        此为系统邮件，请勿直接回复。<br>
        © 农产品融销一体平台 版权所有
      </p>
    </div>
  </div>
</div>
      `,
  });
};
