package matchessrv

func (srv *matchesService) GetTeamsBySeason(season int) ([]string, error) {
	return srv.repository.GetTeamsBySeason(season)
}
