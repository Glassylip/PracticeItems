#!/usr/bin/env python
# coding: utf-8

# In[47]:


import requests
import pandas as pd
from bs4 import BeautifulSoup as bs
import re
import aiohttp
import asyncio
#import nest_asyncio


# In[3]:

# 全国各エリアのURL
regionURLs = {"hokkaido": "https://www.tripadvisor.jp/Attractions-g298143-Activities-c47-Hokkaido.html",
        "aomori": "https://www.tripadvisor.jp/Attractions-g298238-Activities-c47-Akita_Prefecture_Tohoku.html",
        "iwate": "https://www.tripadvisor.jp/Attractions-g298246-Activities-c47-Iwate_Prefecture_Tohoku.html",
        "akita": "https://www.tripadvisor.jp/Attractions-g298238-Activities-c47-Akita_Prefecture_Tohoku.html",
        "miyagi": "https://www.tripadvisor.jp/Attractions-g298248-Activities-c47-Miyagi_Prefecture_Tohoku.html",
        "yamagata": "https://www.tripadvisor.jp/Attractions-g298250-Activities-c47-Yamagata_Prefecture_Tohoku.html",
        "fukushima": "https://www.tripadvisor.jp/Attractions-g298243-Activities-c47-Fukushima_Prefecture_Tohoku.html",
        "ibaraki": "https://www.tripadvisor.jp/Attractions-g298165-Activities-c47-Ibaraki_Prefecture_Kanto.html",
        "tochigi": "https://www.tripadvisor.jp/Attractions-g298181-Activities-c47-Tochigi_Prefecture_Kanto.html",
        "gunma": "https://www.tripadvisor.jp/Attractions-g298163-Activities-c47-Gunma_Prefecture_Kanto.html",
        "saitama": "https://www.tripadvisor.jp/Attractions-g298175-Activities-c47-Saitama_Prefecture_Kanto.html",
        "chiba": "https://www.tripadvisor.jp/Attractions-g298157-Activities-c47-Chiba_Prefecture_Kanto.html",
        "tokyo": "https://www.tripadvisor.jp/Attractions-g1023181-Activities-c47-Tokyo_Prefecture_Kanto.html",
        "kanagawa": "https://www.tripadvisor.jp/Attractions-g298168-Activities-c47-Kanagawa_Prefecture_Kanto.html",
        "niigata": "https://www.tripadvisor.jp/Attractions-g298119-Activities-c47-Niigata_Prefecture_Koshinetsu_Chubu.html",
        "toyama": "https://www.tripadvisor.jp/Attractions-g298125-Activities-c47-Toyama_Prefecture_Hokuriku_Chubu.html",
        "ishikawa": "https://www.tripadvisor.jp/Attractions-g298114-Activities-c47-Ishikawa_Prefecture_Hokuriku_Chubu.html",
        "fukui": "https://www.tripadvisor.jp/Attractions-g298109-Activities-c47-Fukui_Prefecture_Hokuriku_Chubu.html",
        "yamanashi": "https://www.tripadvisor.jp/Attractions-g298127-Activities-c47-Yamanashi_Prefecture_Koshinetsu_Chubu.html",
        "nagano": "https://www.tripadvisor.jp/Attractions-g298117-Activities-c47-Nagano_Prefecture_Koshinetsu_Chubu.html",
        "gifu": "https://www.tripadvisor.jp/Attractions-g298111-Activities-c47-Gifu_Prefecture_Tokai_Chubu.html",
        "shizuoka": "https://www.tripadvisor.jp/Attractions-g298121-Activities-c47-Shizuoka_Prefecture_Tokai_Chubu.html",
        "aichi": "https://www.tripadvisor.jp/Attractions-g298103-Activities-c47-Aichi_Prefecture_Tokai_Chubu.html",
        "mie": "https://www.tripadvisor.jp/Attractions-g298193-Activities-c47-Mie_Prefecture_Tokai_Chubu.html",
        "shiga": "https://www.tripadvisor.jp/Attractions-g298201-Activities-c47-Shiga_Prefecture_Kinki.html",
        "kyoto": "https://www.tripadvisor.jp/Attractions-g298563-Activities-c47-Kyoto_Prefecture_Kinki.html",
        "osaka": "https://www.tripadvisor.jp/Attractions-g298199-Activities-c47-Osaka_Prefecture_Kinki.html",
        "hyogo": "https://www.tripadvisor.jp/Attractions-g298190-Activities-c47-Hyogo_Prefecture_Kinki.html",
        "nara": "https://www.tripadvisor.jp/Attractions-g298197-Activities-c47-Nara_Prefecture_Kinki.html",
        "wakayama": "https://www.tripadvisor.jp/Attractions-g298203-Activities-c47-Wakayama_Prefecture_Kinki.html",
        "tottori": "https://www.tripadvisor.jp/Attractions-g298137-Activities-c47-Tottori_Prefecture_Chugoku.html",
        "shimane": "https://www.tripadvisor.jp/Attractions-g298135-Activities-c47-Shimane_Prefecture_Chugoku.html",
        "okayama": "https://www.tripadvisor.jp/Attractions-g298132-Activities-c47-Okayama_Prefecture_Chugoku.html",
        "hiroshima": "https://www.tripadvisor.jp/Attractions-g298130-Activities-c47-Hiroshima_Prefecture_Chugoku.html",
        "yamaguchi": "https://www.tripadvisor.jp/Attractions-g298140-Activities-c47-Yamaguchi_Prefecture_Chugoku.html",
        "tokushima": "https://www.tripadvisor.jp/Attractions-g676248-Activities-c47-Tokushima_Prefecture_Shikoku.html",
        "kagawa": "https://www.tripadvisor.jp/Attractions-g298231-Activities-c47-Kagawa_Prefecture_Shikoku.html",
        "ehime": "https://www.tripadvisor.jp/Attractions-g298229-Activities-c47-Ehime_Prefecture_Shikoku.html",
        "kochi": "https://www.tripadvisor.jp/Attractions-g298233-Activities-c47-Kochi_Prefecture_Shikoku.html",
        "fukuoka": "https://www.tripadvisor.jp/Attractions-g298206-Activities-c47-Fukuoka_Prefecture_Kyushu.html",
        "saga": "https://www.tripadvisor.jp/Attractions-g298226-Activities-c47-Saga_Prefecture_Kyushu.html",
        "nagasaki": "https://www.tripadvisor.jp/Attractions-g298216-Activities-c47-Nagasaki_Prefecture_Kyushu.html",
        "kumamoto": "https://www.tripadvisor.jp/Attractions-g298212-Activities-c47-Kumamoto_Prefecture_Kyushu.html",
        "oita": "https://www.tripadvisor.jp/Attractions-g298218-Activities-c47-Oita_Prefecture_Kyushu.html",
        "miyazaki": "https://www.tripadvisor.jp/Attractions-g298214-Activities-c47-Miyazaki_Prefecture_Kyushu.html",
        "kagoshima": "https://www.tripadvisor.jp/Attractions-g298209-Activities-c47-Kagoshima_Prefecture_Kyushu.html",
        "okinawa": "https://www.tripadvisor.jp/Attractions-g298221-Activities-c47-Okinawa_Prefecture.html"
       }


# In[4]:


def get_mainURL(region):
    url = regionURLs[region]
    html_ = requests.get(url)
    soup = bs(html_.text,'lxml')
    html = [ "https://www.tripadvisor.jp"+re.findall(r"href=\"(.+)\" onclick",str(i))[0]  for i in soup.select('div[class="shelf_title_container"]>a') ]
    
    
    html_ = requests.get(html[0])
    soup = bs(html_.text,'lxml')
    for i in soup.select('div[class="shelf_title_container"]>a'):
        target = "https://www.tripadvisor.jp"+re.findall(r"href=\"(.+)\" onclick",str(i))[0]
        if target not in html and "#" not in target:
            html.append(target)
    return html


# In[48]:


#nest_asyncio.apply()


# In[54]:

# 非同期処理スクレイピング
async def fetch(session, url):
    async with session.get(url) as html:
        return await html.text()


# In[55]:


async def parser(html):
    soup = bs(html,'lxml')
    urls = soup.select('div[class="listing_title title_with_snippets"] > a')
    url = re.findall(r"href=\"(.+)\" onclick",str(urls))
    spot_urls.extend( ["https://www.tripadvisor.jp/"+i for i in url ] )


# In[51]:


async def download(url):
    async with aiohttp.ClientSession() as session:
        html = await fetch(session, url) 
        await parser(html)


# In[56]:
# 全国観光地URL取得
print("start")
region = list(regionURLs.keys())
for i in range(len(regionURLs.keys())):
    main_urls = get_mainURL(region[i])
    spot_urls = []
    loop = asyncio.get_event_loop()
    tasks = [asyncio.ensure_future(download(main_url)) for main_url in main_urls ]
    tasks = asyncio.gather(*tasks)
    loop.run_until_complete(tasks)
    spot_urls=list(set(spot_urls))
    if i!= 0:
        df1 = pd.DataFrame(spot_urls,columns=[region[i]])
        df = pd.concat([df,df1],axis=1)
    else:
        df = pd.DataFrame(spot_urls,columns=[region[i]])


# In[58]:

# データ保存
df.to_pickle("allURL.pkl")
print("success")

# In[ ]:




