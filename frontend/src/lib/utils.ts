import { RANKS } from "$lib/constants";
import type { main } from "$wails/go/models";
import { GetRegistryData } from "$wails/go/main/App";

export async function getRegistryData()
{
    let data = await GetRegistryData();

    const leagueData = convertTrophiesToLeague(data.userTrophies);
    if (leagueData)
        data = {
            ...data,
            userLeague: leagueData.league,
            userDivision: leagueData.division,
            userStars: leagueData.stars
        };

    return data;
}

export function isRegistryDataValid(data: main.RegistryData)
{
    return data.userId !== "" && data.username !== "";
}

export function convertTrophiesToLeague(trophies: number)
{
    const result = RANKS.filter(rank => trophies >= rank.minTrophy).pop();

    return result === undefined ? null : {
        league: result.league,
        division: result.division,
        stars: trophies - result.minTrophy
    };
}