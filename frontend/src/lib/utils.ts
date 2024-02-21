import { RANKS } from "$lib/constants";
import type { UserData } from "$lib/types";
import { GetUserData } from "$wails/go/main/App";

export async function getUserData()
{
    const base = await GetUserData();

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
    } as UserData;

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

export function isUserDataValid(data: UserData)
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