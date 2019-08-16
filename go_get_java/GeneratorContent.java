import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.SecureRandom;

import java.util.Base64;
import java.util.Base64.Encoder;
import java.util.Base64.Decoder;
import java.security.NoSuchAlgorithmException;
import java.util.Arrays;

import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;

public class GeneratorKey {
    private static final String ALGORITHM = "AES/CBC/PKCS5Padding";

    public static String getImageStr(String imgFile) {
        InputStream inputStream = null;
        byte[] data = null;
        try {
            inputStream = new FileInputStream(imgFile);
            data = new byte[inputStream.available()];
            inputStream.read(data);
            inputStream.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
        // 加密
        Encoder encoder = Base64.getEncoder();
        return encoder.encodeToString(data);
    }

    /**
     * aes解密-256位
     */
    public static String aesDecrypt(String encryptContent, String password) {
        if (password.length() !=16) {
            System.out.println("password must be is 16 bytes");
        }
        try {
            Decoder decoder = Base64.getDecoder();
            KeyGenerator keyGen = KeyGenerator.getInstance("AES");
            SecureRandom secureRandom = SecureRandom.getInstance("SHA1PRNG");
            secureRandom.setSeed(password.getBytes());
            keyGen.init(256, secureRandom);
            SecretKey secretKey = keyGen.generateKey();
            byte[] enCodeFormat = secretKey.getEncoded();
            SecretKeySpec key = new SecretKeySpec(enCodeFormat, "AES");
            Cipher cipher = Cipher.getInstance(ALGORITHM);
            IvParameterSpec iv = new IvParameterSpec(password.getBytes());//使用CBC模式，需要一个向量iv，可增加加密算法的强度
            cipher.init(Cipher.DECRYPT_MODE, key,iv);
            return new String(cipher.doFinal(decoder.decode(encryptContent)));
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }

    /**
     * aes加密-256位
     */
    public static String aesEncrypt(String content, String password) {
        if (password.length() !=32) {
            System.out.println("password must be is 16 bytes");
        }
        try {
            Encoder encoder = Base64.getEncoder();
            byte[] enCodeFormat = secretKey.getEncoded();
            SecretKeySpec key = new SecretKeySpec(enCodeFormat, "AES");
            Cipher cipher = Cipher.getInstance(ALGORITHM);
            IvParameterSpec iv = new IvParameterSpec(password.getBytes());//使用CBC模式，需要一个向量iv，可增加加密算法的强度
            cipher.init(Cipher.ENCRYPT_MODE, key,iv);
            return encoder.encodeToString(cipher.doFinal(content.getBytes("UTF-8")));
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }

	public static void main(String[] args) {
        String data = args[0];
        String type = args[1];
        // 16个明文字符串:Eg encryptedKey1234
		String password = args[2];
        if (type.equals("e")) {
            String encryptData = aesEncrypt(data, password);
            System.out.println("加密结果："+encryptData);
        } else if (type.equals("d")) {
            String decryptData = aesDecrypt(data, password);
            System.out.println("解密结果："+decryptData);
        } else {
            System.out.println("操作方式错误!");
        }
	}
}
