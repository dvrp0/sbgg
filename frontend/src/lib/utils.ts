import type { UserData } from "$lib/types";

export function isUserDataValid(data: UserData)
{
    return data.UserId !== "" && data.UserName !== "";
}