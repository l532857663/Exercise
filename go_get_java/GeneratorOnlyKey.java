import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.SecureRandom;

import java.util.Base64;
import java.util.Base64.Encoder;
import java.security.NoSuchAlgorithmException;
import javax.crypto.Cipher;
import java.util.Arrays;

public class GeneratorOnlyKey {
    private static final String ALGORITHM = "AES/CBC/PKCS5Padding";
    /**
     * aes加密-256位
     */
    public static String[] aesEncrypt(String password) {
        if (password.length() !=16) {
            System.out.println("password must be is 16 bytes");
        }
        try {
            String[] encryptData = new String[2];
            Encoder encoder = Base64.getEncoder();
            KeyGenerator keyGen = KeyGenerator.getInstance("AES");
            SecureRandom secureRandom = SecureRandom.getInstance("SHA1PRNG");
            secureRandom.setSeed(password.getBytes());
            keyGen.init(256, secureRandom);
            SecretKey secretKey = keyGen.generateKey();
            byte[] enCodeFormat = secretKey.getEncoded();
            SecretKeySpec key = new SecretKeySpec(enCodeFormat, "AES");
            String keyS = encoder.encodeToString(enCodeFormat);
            encryptData[0] = keyS;
            byte[] KeyB = key.getEncoded();
            String KeySB = new String(KeyB);
            System.out.println(Arrays.toString(KeyB));
            String keyBS = encoder.encodeToString(KeyB);
            System.out.println("加解密keyS："+KeySB);
            System.out.println("加解密key："+keyBS);

            IvParameterSpec iv = new IvParameterSpec(password.getBytes());//使用CBC模式，需要一个向量iv，可增加加密算法的强度
            byte[] ivB = iv.getIV();

            Cipher cipher = Cipher.getInstance(ALGORITHM);
            cipher.init(Cipher.ENCRYPT_MODE, key,iv);
            byte[] ivC = cipher.getIV();
            System.out.println(Arrays.toString(ivC));

            String ivS = encoder.encodeToString(ivB);
            System.out.println("加解密iv："+ivS);

            String ivSC = encoder.encodeToString(ivC);
            System.out.println("加解密iv："+ivSC);
            return encryptData;
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }

	public static void main(String[] args) {
        // 16个明文字符串:Eg encryptedKey1234
		String password = args[0];
        String[] encryptData = aesEncrypt(password);
        System.out.println("加解密key："+encryptData[0]);
        System.out.println(encryptData);
	}
}
