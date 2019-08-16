public class GeneratorRsa {
    /**
     * 对内容进行加密
     *
     * @param content
     * @param privateKey 私钥
     * @return
     */
    public static String encrypt(String content, String privateKey) {
        PrivateKey pk = getPrivateKey(privateKey);
        byte[] data = encryptByPrivateKey(content, pk);
        String res = null;
        try {
            res = Base64.encode(data);
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        return res;

    }

    /**
     * 得到私钥对象
     * @param privateKey 密钥字符串（经过base64编码的秘钥字节）
     * @throws Exception
     */
    public static PrivateKey getPrivateKey(String privateKey)  {
        try {
            byte[] keyBytes;

            keyBytes = Base64.decode(privateKey);

            PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);

            KeyFactory keyFactory = KeyFactory.getInstance("RSA");

            PrivateKey privatekey = keyFactory.generatePrivate(keySpec);

            return privatekey;
        }catch(Exception e){
            e.printStackTrace();
        }
        return null;
    }

    /**
     * 通过私钥加密
     *
     * @param content
     * @param pk
     * @return,加密数据，未进行base64进行加密
     */
    protected static byte[] encryptByPrivateKey(String content, PrivateKey pk) {

        try {
            Cipher ch = Cipher.getInstance(ALGORITHM);
            ch.init(Cipher.ENCRYPT_MODE, pk);
            return ch.doFinal(content.getBytes(CHAR_SET));
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("通过私钥加密出错");
        }
        return null;

    }
}
