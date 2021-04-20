#!/usr/bin/python3

from urllib import parse

import urllib.request
import json, time
import sys

UrlList = [
    # MDEX swap的交易通证 + 时间戳
    'https://ht.mdex.com/tokenlist.json?t=1611745257666',
    # Ark Swap的交易通证
    'https://arkswap.org/tokenLists/ArkSwap.json',
    ]

# 常量
headers ={
    "User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36",
}
# 文件名
TokenSymbolLogName = './token_symbol.log'
SqlLogName = './sql.log'
IconFile = '/home/w123/icon/'
## Heco主网信息
HecoID = 128
HecoHost = 'https://http-mainnet-node.huobichain.com'
HecoTokenDetailsUrl = 'https://api-scan.hecochain.com/hsc/getTokenDetails/'
## Heco测试网信息
HecoTestID = 256
HecoTestHost = 'https://http-testnet.hecochain.com'
HecoTestTokenDetailsUrl = 'https://api-scan-testnet.hecochain.com/hsc/getTokenDetails/'
# 0: token_symbol 1: token_type 2: token_contract 3: token_name 4: decimals 5: icon_url 6: description 7: default_gas_limit
SqlInsert = '''
INSERT INTO `bingoo_db`.`bingoo_token`(`created_at`, `updated_at`, `deleted_at`, `system_id`, `platform_id`, `platform_coin_symbol`, `token_symbol`, `token_type`, `token_standard`, `token_contract`, `token_full_name`, `decimals`, `show_decimals`, `show_decimals_with_no_places`, `icon_url`, `background_image_url`, `description`, `transfer_action`, `default_gas_limit`, `state`, `version_compatibility_state`, `is_default`, `is_default_view`, `is_major`, `is_show`, `is_standard_show`, `is_full_name_show`) VALUES (CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), NULL, 2, 1010, 'HT', '{0}', {1}, 'HRC20', '{2}', '{3}', {4}, 6, 2, '{5}', '', '{6}', '', {7}, '1', '1', '0', '1', '0', '0', '0', '0');
'''

def callback(a1,a2,a3):
    """
        @a1:目前为此传递的数据块数量
        @a2:每个数据块的大小，单位是byte,字节
        @a3:远程文件的大小
    """
    download_pg = 100.0*a1*a2/a3
    if download_pg > 100:
        download_pg = 100
    print("%.2f%%" %download_pg,)

def urlGet(url):
    # url 作为Request()方法的参数，构造并返回一个Request对象
    request = urllib.request.Request(url,headers=headers)
    # Request对象作为urlopen()方法的参数，发送给服务器并接收响应
    response = urllib.request.urlopen(request)
    html = response.read().decode('utf-8')
    # 转化为json结构
    resJson = json.loads(html)
    return resJson

def logTokenInfo(tokenInfo, chainId, tokenSymbolFile, sqlFile):
    iconUrl = 'https://res.bingooglobal.com/'
    if chainId == HecoTestID:
        url = HecoTestTokenDetailsUrl + tokenInfo['address']
        iconUrl += 'dev/token-icon/'
    elif chainId == HecoID:
        url = HecoTokenDetailsUrl + tokenInfo['address']
        iconUrl += 'prod/token-icon/'
    else:
        return False

    resJson = urlGet(url)
    tokenData = resJson['data']
    # 判断token信息
    if resJson['status'] != 1 or resJson['message'] != "select success":
        return False
    print("tokenInfo:", tokenInfo)
    print("tokenData:", tokenData)
    iconName = tokenData['symbol']+'-HRC20_3V.png'
    iconUrl += iconName
    tokenAll = {
            "token_symbol": tokenData['symbol'],
            "token_type": 0,
            "token_contract": tokenInfo['address'],
            "token_name": tokenInfo['name'],
            "decimals": tokenData['decimals'],
            "icon_url": iconUrl,
            "description": "",
            "default_gas_limit": 60000,
    }
    # 生成sql
    sqlData = SqlInsert.format(tokenAll["token_symbol"],tokenAll["token_type"],tokenAll["token_contract"],tokenAll["token_name"],tokenAll["decimals"],tokenAll["icon_url"],tokenAll["description"],tokenAll["default_gas_limit"])
    # 保存Token
    tokenSymbolFile.write(tokenData['symbol']+'_'+str(chainId)+"\n")
    # 保存sql
    sqlFile.write(sqlData)
    # 保存Icon
    try:
        pass
        # urllib.request.urlretrieve(tokenInfo['logoURI'], IconFile+iconName, callback)
    except urllib.error.URLError as e:
        print("Can't get the %s token icon: %s" % (tokenData['symbol'], tokenInfo['logoURI']))
        print("error:", e)
    return True

def filterTokenMap(tokenMap):
    # token_symbol记录
    tokeSymbolLog = open(TokenSymbolLogName, 'a+')
    sqlLog = open(SqlLogName, 'a+')
    for tokenInfo in tokenMap:
        # Heco网段判断
        tokenSymbol = tokenInfo['symbol']
        chainId = tokenInfo['chainId']
        if chainId == HecoID or chainId == HecoTestID:
            ok = logTokenInfo(tokenInfo, chainId, tokeSymbolLog, sqlLog)
            if ok == False :
                print("logTokenInfo error, symbol: %s, chainId: %d" % (tokenSymbol, chainId))
                continue
    # 关闭文件
    tokeSymbolLog.close()
    sqlLog.close()

def main():
    print("Start")
    for url in UrlList:
        resJson = urlGet(url)
        if len(resJson['tokens']) > 0:
            tokenMap = resJson['tokens']

        # 过滤获取到的token
        filterTokenMap(tokenMap)
    print("End")

if __name__ == '__main__':
    main()
