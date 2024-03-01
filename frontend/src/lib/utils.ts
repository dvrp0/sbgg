import { RANKS } from "$lib/constants";
import type { RegistryData } from "$lib/types";
import { GetRegistryData } from "$wails/go/main/App";

export async function getUserData()
{
    const base = await GetRegistryData();

    let data = {
        userId: base.UserId,
        username: base.Username,
        userTrophies: base.UserTrophies,
        userRank: base.UserRank,
        userLevel: base.UserLevel,
        timeMatchmakingStarted: base.TimeMatchmakingStarted,
        gameTurns: base.GameTurns,
        timeMatchStarted: base.TimeMatchStarted,
        rankedPlayed: base.RankedPlayed,
        rankedWon: base.RankedWon
    } as RegistryData;

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

export function isRegistryDataValid(data: RegistryData)
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