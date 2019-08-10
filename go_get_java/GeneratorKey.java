import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import java.security.SecureRandom;
import java.util.Base64;
import java.util.Base64.Encoder;
import java.security.NoSuchAlgorithmException;

public class GeneratorKey {
	public static String createEncryptOrDecryptKey(String password) {
		KeyGenerator keyGen = null;
		try {
			keyGen = KeyGenerator.getInstance("AES");
			SecureRandom secureRandom = SecureRandom.getInstance("SHA1PRNG");
			secureRandom.setSeed(password.getBytes());
			keyGen.init(256, secureRandom);
			SecretKey secretKey = keyGen.generateKey();
			byte[] enCodeFormat = secretKey.getEncoded();
			Encoder encoder = Base64.getEncoder();
			String key = encoder.encodeToString(enCodeFormat);
			return key;
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		}
		return null;
	}

	public static void main(String[] args) {
		String password = args[0];
		String encryptOrDecryptKey = createEncryptOrDecryptKey(password);
		System.out.println("加解密密码串Key："+encryptOrDecryptKey);
	}
}
