#!/usr/bin/env python
# coding: utf-8

# In[ ]:

from requests_html import AsyncHTMLSession,user_agent
import os,zipfile
from fake_useragent import UserAgent
from re import findall
import pandas as pd
from numpy import nan
import asyncio
from tqdm import tqdm
from shutil import rmtree

# get tourist destination data
async def spot_data(r,i):
    find = r.html.find
    spot_name[i].append(find('h1#HEADING')[0].text)
    try:
        spot_score[i].append(int(findall(r"bubble_(\d\d)",str(find('div._1NKYRldB>span')[0]))[0]) / 10)
    except:
        spot_score[i].append(nan)
    try:
        spot_tag[i].extend( [i.text for i in find('span>a._1cn4vjE4')] )
    except:
        spot_tag[i].append(nan)
    try:
        count[i].append(find('div._1NKYRldB')[0].text)
    except:
        count[i].append(nan)
    try:
        latlng = str(find("img._384yV20z"))
        lat[i].append(findall(r"\d+\.\d+\,\d+\.\d+",latlng)[0].split(",")[1])
        lng[i].append(findall(r"\d+\.\d+\,\d+\.\d+",latlng)[0].split(",")[0])
    except:
        lat[i].append(nan)
        lng[i].append(nan)
    try:
        ranking[i].append(find('div.eQSJNhO6')[0].text)
    except:
        ranking[i].append(nan)
    try:
        address[i].append(find('div.LjCWTZdN > span')[1].text)
    except:
        address[i].append(nan)
    try:
        tel[i].append(find('div.LjCWTZdN > div')[1].text)
    except:
        tel[i].append(nan)

# save data
def datasave(path):
    # データ整形
    t = [[] for i in range(len(spot_tag))]
    for i in range(len(spot_tag)):
        t[i].append(",".join(spot_tag[i]))

    spot_name1 = pd.DataFrame(spot_name)
    spot_score1 = pd.DataFrame(spot_score)
    spot_tag1 = pd.DataFrame(t)
    count1 = pd.DataFrame(count)
    lat1 = pd.DataFrame(lat)
    lng1 = pd.DataFrame(lng)
    tel1 = pd.DataFrame(tel)
    ranking1 = pd.DataFrame(ranking)
    address1 = pd.DataFrame(address)

    spot_urls1 = pd.DataFrame(spot_urls)


    df_spot = pd.concat([spot_name1, spot_score1, spot_tag1, count1, ranking1, address1, tel1, lat1, lng1, spot_urls1], axis=1)
    df_spot.columns = ["SpotName", "Score", "Tag", "Count", "Ranking", "Address", "Tel", "Lat", "Lng", "URL"]
    df_spot = df_spot.dropna(axis=0, thresh=7).reset_index(drop=True)
    df_spot.to_csv(path + "/" + region + "_spot.csv",encoding="utf_8_sig",index=False)

# Parallel processing
async def do(semaphore):
    async def getURL(url,i):
        async with semaphore:
            flag = 0
            while True:
                try:
                    headers = {
                    "User-Agent":ua.random
                    }

                    r = await asession.get(url,headers=headers)
                    await r.html.arender(wait=1,sleep=1)
                    await spot_data(r,i)
                    break
                except:
                    flag += 1
                    spot_name[i].clear()
                    spot_score[i].clear()
                    count[i].clear()
                    lat[i].clear()
                    lng[i].clear()
                    ranking[i].clear()
                    spot_tag[i].clear()
                    tel[i].clear()
                    address[i].clear()
                    if flag > 1:
                        break
    tasks = [getURL(spot_urls[i], i) for i in range(len(spot_urls))]
    [ await f for f in tqdm(asyncio.as_completed(tasks), total=len(tasks)) ]
"""
def getURL(url,i):
    session = HTMLSession()
    headers = {
    "User-Agent":ua.random
    }
    r = session.get(url,headers=headers)
    r.html.render(timeout = 20)
    spot_data(r,i)
    session.close()
"""
    
if __name__ == "__main__":
    
    # Get the regional url
    while True:
        region = input("県名(例：saitama)を入力してください： ")
        df_url = pd.read_pickle("allURL.pkl")
        if region in df_url.columns or region== "all":
            break
        else:
            print("県名が違います")
    if region == "all":
        length = len(df_url.columns)
    else:
        length = 1
        
    # UserAgent
    ua = UserAgent()

    # CPU core
    semaphore = asyncio.Semaphore(value=8)

    asession = AsyncHTMLSession()
    
    # File save destination
    path = "spot_data"
    
    # Create destination folder
    try:  
        os.mkdir("spot_data")
    except:
        pass
    
    for i in range(length):
        if region == "all":
            spot_urls = list(df_url[df.columns[i]].dropna())
            print(df.columns[i]," start")
        else:
            spot_urls = list(df_url[region].dropna())
        #spot data
        spot_name = [ [] for i in range(len(spot_urls)) ]
        spot_score = [ [] for i in range(len(spot_urls)) ]
        spot_tag = [ [] for i in range(len(spot_urls)) ]
        count = [ [] for i in range(len(spot_urls)) ]
        lat = [ [] for i in range(len(spot_urls)) ]
        lng = [ [] for i in range(len(spot_urls)) ]
        ranking = [ [] for i in range(len(spot_urls)) ]
        tel = [ [] for i in range(len(spot_urls)) ]
        address = [ [] for i in range(len(spot_urls)) ]

        
        # Scraping
        loop = asyncio.get_event_loop()
        # tasks = [asyncio.ensure_future(getURL(spot_urls[i], i)) for i in range(len(spot_urls))]
        # tasks = asyncio.gather(*tasks)
        loop.run_until_complete(do(semaphore))
        datasave(path)
        """
        for i in tqdm(range(len(spot_urls))):
            getURL(spot_urls[i],i)
        """

    
     # Data compression
    if len([lists for lists in os.listdir(path) if os.path.isfile(os.path.join(path, lists))]) == 47:
        startdir = path 
        file_news = startdir +'.zip' # zip file name
        z = zipfile.ZipFile(file_news,'w',zipfile.ZIP_DEFLATED) 
        for dirpath, dirnames, filenames in os.walk(startdir):
            fpath = dirpath.replace(startdir,'')
            fpath = fpath and fpath + os.sep or ''
            for filename in filenames:
                z.write(os.path.join(dirpath, filename),fpath+filename)
        z.close()
        rmtree(startdir)
      
    print("end")

