package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
)

//Used for queries to anilist
//Anilist uses graphql, genqlient (github.com/Khan/genqlient/graphql) is used to
//query for data
//Queries written in genqlient.graphql, functions generated with "go run github.com/Khan/genqlient --init"

func getUId(username string) int {
	ctx := context.Background()
	client := graphql.NewClient("https://graphql.anilist.co", http.DefaultClient)
	resp, err := queryUid(ctx, client, username)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(strconv.Itoa(resp.User.Id))
	return resp.User.Id
}

func getRecentActv(username string) {
	ctx := context.Background()
	client := graphql.NewClient("https://graphql.anilist.co", http.DefaultClient)
	resp, err := queryRecentActv(ctx, client, getUId(username), int(time.Now().Unix()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func getFavAnime(username string) []queryFavAnimeUserFavouritesAnimeMediaConnectionNodesMedia {
	ctx := context.Background()
	client := graphql.NewClient("https://graphql.anilist.co", http.DefaultClient)
	resp, err := queryFavAnime(ctx, client, getUId(username))
	if err != nil {
		log.Fatal(err)
	}

	//marshal, _ := json.MarshalIndent(resp, "", "	")
	//fmt.Println(string(marshal))
	nodes := resp.User.Favourites.Anime.Nodes
	for i := range nodes {
		fmt.Printf("%d. %s\n", i+1, nodes[i].Title.Romaji)
	}
	return nodes
	//fmt.Println(resp.User.Favourites.Anime)
}

func getFavManga(username string) []queryFavMangaUserFavouritesMangaMediaConnectionNodesMedia {
	ctx := context.Background()
	client := graphql.NewClient("https://graphql.anilist.co", http.DefaultClient)
	resp, err := queryFavManga(ctx, client, getUId(username))
	if err != nil {
		log.Fatal(err)
	}

	nodes := resp.User.Favourites.Manga.Nodes
	for i := range nodes {
		fmt.Printf("%d. %s\n", i+1, nodes[i].Title.Romaji)
	}

	return nodes
}
func getFavChara(username string) []queryFavCharaUserFavouritesCharactersCharacterConnectionNodesCharacter {
	ctx := context.Background()
	client := graphql.NewClient("https://graphql.anilist.co", http.DefaultClient)
	resp, err := queryFavChara(ctx, client, getUId(username))
	if err != nil {
		log.Fatal(err)
	}

	nodes := resp.User.Favourites.Characters.Nodes
	for i := range nodes {
		if nodes[i].Name.Last == "" {
			fmt.Printf("%d. %s\n", i+1, nodes[i].Name.First)
		} else if nodes[i].Name.First == "" {
			fmt.Printf("%d. %s\n", i+1, nodes[i].Name.Last)
		} else {
			fmt.Printf("%d. %s, %s\n", i+1, nodes[i].Name.Last, nodes[i].Name.First)
		}

	}
	return nodes
}
